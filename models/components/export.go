// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Export struct {
	ID           *string       `json:"id,omitempty"`
	Params       *ExportCreate `json:"params,omitempty"`
	Volumename   *string       `json:"volumename,omitempty"`
	Snapshotname *string       `json:"snapshotname,omitempty"`
	State        *string       `json:"state,omitempty"`
	Status       *string       `json:"status,omitempty"`
	Progress     *int64        `json:"progress,omitempty"`
	Xqn          *string       `json:"xqn,omitempty"`
	Wwn          *string       `json:"wwn,omitempty"`
}

func (o *Export) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Export) GetParams() *ExportCreate {
	if o == nil {
		return nil
	}
	return o.Params
}

func (o *Export) GetVolumename() *string {
	if o == nil {
		return nil
	}
	return o.Volumename
}

func (o *Export) GetSnapshotname() *string {
	if o == nil {
		return nil
	}
	return o.Snapshotname
}

func (o *Export) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *Export) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Export) GetProgress() *int64 {
	if o == nil {
		return nil
	}
	return o.Progress
}

func (o *Export) GetXqn() *string {
	if o == nil {
		return nil
	}
	return o.Xqn
}

func (o *Export) GetWwn() *string {
	if o == nil {
		return nil
	}
	return o.Wwn
}
