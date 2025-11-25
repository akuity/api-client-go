package http

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"k8s.io/utils/env"
)

var maxCookieLength = 4093

func init() {
	if val := strings.TrimSpace(env.GetString("MAX_COOKIE_LENGTH", "")); val != "" {
		if length, err := strconv.Atoi(val); err != nil {
			panic(fmt.Errorf("failed to parse MAX_COOKIE_LENGTH: %v", err))
		} else {
			// 100 characters is required to store attributes so smaller values does not make sense
			if maxCookieLength < 100 {
				panic(fmt.Errorf("MAX_COOKIE_LENGTH must be at least 100"))
			}
			maxCookieLength = length
		}
	}
}

func SetAKPTokenCookie(w http.ResponseWriter, token string, secure bool, domain string) {
	flags := []string{"path=/", "SameSite=lax", "httpOnly"}
	if secure {
		flags = append(flags, "Secure")
	}
	if domain != "" {
		flags = append(flags, "domain="+domain)
	}
	cookies := splitCookie(TokenCookieName, token, strings.Join(flags, "; "))
	for _, cookie := range cookies {
		w.Header().Add("Set-Cookie", cookie)
	}
}

func SetAimsTokenCookie(w http.ResponseWriter, token string, secure bool, domain string) {
	flags := []string{"path=/", "SameSite=lax", "httpOnly"}
	if secure {
		flags = append(flags, "Secure")
	}
	if domain != "" {
		flags = append(flags, "domain="+domain)
	}
	cookies := splitCookie(AimsTokenCookieName, token, strings.Join(flags, "; "))
	for _, cookie := range cookies {
		w.Header().Add("Set-Cookie", cookie)
	}
}

func GetAKPTokenCookie(r *http.Request) (string, error) {
	return joinCookies(TokenCookieName, r.Cookies())
}

func GetAimsTokenCookie(r *http.Request) (string, error) {
	return joinCookies(AimsTokenCookieName, r.Cookies())
}

func GetArgoCDTokenCookie(r *http.Request) (string, error) {
	return joinCookies(ArgoCDTokenCookieName, r.Cookies())
}

func SetArgoCDRequestTokenCookie(token string, httpOnly, secure bool, domain string) ([]*http.Cookie, error) {
	cookies := splitCookie(ArgoCDTokenCookieName, token, "")
	httpCookies := []*http.Cookie{}
	for _, cookie := range cookies {
		splits := strings.Split(cookie, "=")
		if len(splits) != 2 {
			return nil, fmt.Errorf("invalid cookie format recieved from splitCookie: %s", cookie)
		}
		httpCookies = append(httpCookies, &http.Cookie{
			Name:     splits[0],
			Value:    splits[1],
			HttpOnly: httpOnly,
			Secure:   secure,
			Domain:   domain,
		})
	}
	return httpCookies, nil
}

// Functions below are borrowed from
// https://github.com/argoproj/argo-cd/blob/61abc805731962cfd20159c315d8ed436870b0e7/util/http/http.go#L36

// maxCookieValueLength returns the maximum length of cookie value based on key and attributes
func maxCookieValueLength(key, attributes string) int {
	if len(attributes) > 0 {
		return maxCookieLength - (len(key) + 3) - (len(attributes) + 2)
	}
	return maxCookieLength - (len(key) + 3)
}

// browser has limit on size of cookie, currently 4kb. In order to
// support cookies longer than 4kb, we split cookie into multiple 4kb chunks.
// first chunk will be of format token=<numberOfChunks>:token; attributes
func splitCookie(key, value, attributes string) []string {
	var cookies []string
	valueLength := len(value)
	// cookie: name=value; attributes and key: key-(i) e.g. token-1
	maxValueLength := maxCookieValueLength(key, attributes)
	numberOfChunks := int(math.Ceil(float64(valueLength) / float64(maxValueLength)))

	var end int
	for i, j := 0, 0; i < valueLength; i, j = i+maxValueLength, j+1 {
		end = i + maxValueLength
		if end > valueLength {
			end = valueLength
		}

		var cookie string
		if j == 0 && numberOfChunks == 1 {
			cookie = fmt.Sprintf("%s=%s", key, value[i:end])
		} else if j == 0 {
			cookie = fmt.Sprintf("%s=%d:%s", key, numberOfChunks, value[i:end])
		} else {
			cookie = fmt.Sprintf("%s-%d=%s", key, j, value[i:end])
		}
		if attributes != "" {
			cookie = fmt.Sprintf("%s; %s", cookie, attributes)
		}
		cookies = append(cookies, cookie)
	}
	return cookies
}

// joinCookies combines chunks of cookie based on key as prefix. It returns cookie
// value as string. cookieString is of format key1=value1; key2=value2; key3=value3
// first chunk will be of format token=<numberOfChunks>:token; attributes
func joinCookies(key string, cookieList []*http.Cookie) (string, error) {
	cookies := make(map[string]string)
	for _, cookie := range cookieList {
		if !strings.HasPrefix(cookie.Name, key) {
			continue
		}
		cookies[cookie.Name] = cookie.Value
	}

	var sb strings.Builder
	var numOfChunks int
	var err error
	var token string
	var ok bool

	if token, ok = cookies[key]; !ok {
		return "", fmt.Errorf("failed to retrieve cookie %s", key)
	}
	parts := strings.Split(token, ":")

	if len(parts) == 2 {
		if numOfChunks, err = strconv.Atoi(parts[0]); err != nil {
			return "", err
		}
		sb.WriteString(parts[1])
	} else if len(parts) == 1 {
		numOfChunks = 1
		sb.WriteString(parts[0])
	} else {
		return "", fmt.Errorf("invalid cookie for key %s", key)
	}

	for i := 1; i < numOfChunks; i++ {
		sb.WriteString(cookies[fmt.Sprintf("%s-%d", key, i)])
	}
	return sb.String(), nil
}
