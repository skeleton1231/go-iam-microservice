package genericclioptions

// Defines flag for iamctl.
const (
	FlagIAMConfig     = "iamconfig"
	FlagBearerToken   = "user.token"
	FlagUsername      = "user.username"
	FlagPassword      = "user.password"
	FlagSecretID      = "user.secret-id"
	FlagSecretKey     = "user.secret-key"
	FlagCertFile      = "user.client-certificate"
	FlagKeyFile       = "user.client-key"
	FlagTLSServerName = "server.tls-server-name"
	FlagInsecure      = "server.insecure-skip-tls-verify"
	FlagCAFile        = "server.certificate-authority"
	FlagAPIServer     = "server.address"
	FlagTimeout       = "server.timeout"
	FlagMaxRetries    = "server.max-retries"
	FlagRetryInterval = "server.retry-interval"
)
