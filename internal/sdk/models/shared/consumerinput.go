// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConsumerInput struct {
	// Field for storing an existing unique ID for the Consumer - useful for mapping Kong with users in your existing database. You must send either this field or `username` with the request.
	CustomID *string `json:"custom_id,omitempty"`
	// An optional set of strings associated with the Consumer for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// The unique username of the Consumer. You must send either this field or `custom_id` with the request.
	Username *string `json:"username,omitempty"`
}

func (o *ConsumerInput) GetCustomID() *string {
	if o == nil {
		return nil
	}
	return o.CustomID
}

func (o *ConsumerInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ConsumerInput) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}
