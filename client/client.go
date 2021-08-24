package client

// AWSMFA specifies the
type AWSMFA struct {
	DisabledUsers []string
	count         uint32
}

func getMFADisabled() AWSMFA {
  users := AWSMFA {
	DisabledUsers: []string{"isuru", "siriwardana"},
	count: 2,
  }

  return users
}

// GetMFADisabled is used to
func GetMFADisabled() AWSMFA {
	return getMFADisabled()
}
