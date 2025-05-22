package featuresv1

func (x FeatureStatus) Enabled() bool {
	return x == FeatureStatus_FEATURE_STATUS_ENABLED
}
