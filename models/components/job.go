// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Args struct {
}

type Job struct {
	ID        *int64  `json:"id,omitempty"`
	Type      *string `json:"type,omitempty"`
	Object    *string `json:"object,omitempty"`
	Args      *Args   `json:"args,omitempty"`
	State     *string `json:"state,omitempty"`
	Status    *string `json:"status,omitempty"`
	Progress  *int64  `json:"progress,omitempty"`
	Starttime *string `json:"starttime,omitempty"`
	Endtime   *string `json:"endtime,omitempty"`
	Username  *string `json:"username,omitempty"`
	Useremail *string `json:"useremail,omitempty"`
}

func (o *Job) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Job) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Job) GetObject() *string {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *Job) GetArgs() *Args {
	if o == nil {
		return nil
	}
	return o.Args
}

func (o *Job) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *Job) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Job) GetProgress() *int64 {
	if o == nil {
		return nil
	}
	return o.Progress
}

func (o *Job) GetStarttime() *string {
	if o == nil {
		return nil
	}
	return o.Starttime
}

func (o *Job) GetEndtime() *string {
	if o == nil {
		return nil
	}
	return o.Endtime
}

func (o *Job) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *Job) GetUseremail() *string {
	if o == nil {
		return nil
	}
	return o.Useremail
}