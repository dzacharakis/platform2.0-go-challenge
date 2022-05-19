package model

type Asset interface {
	GetAssetID() int
}

func (v Variable) GetAssetID() int {
	return v.AssetID
}

func (i Insight) GetAssetID() int {
	return i.AssetID
}

func (c Chart) GetAssetID() int {
	return c.AssetID
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Favourite struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id" binding:"required"`
	AssetID int `json:"asset_id" binding:"required"`
}

type Variable struct {
	AssetID        int    `json:"id"`
	AssetType      string `json:"type,omitempty"`
	Name           string `json:"name"`
	VarType        string `json:"var_type"`
	PossibleValues string `json:"possible_values,omitempty"` // setting to sql.NullString creates a problem, another struct with a Valid field
	Unit           string `json:"unit,omitempty"`            // the same
}

type Insight struct {
	AssetID   int    `json:"id"`
	AssetType string `json:"type,omitempty"`
	ChartID   int    `json:"chart_id"`
	Statement string `json:"statement"`
}

type Chart struct {
	AssetID   int    `json:"id"`
	AssetType string `json:"type,omitempty"`
	Title     string `json:"title"`
	XName     string `json:"x_name"`
	YName     string `json:"y_name"`

	Data []ChartData `json:"data,omitempty"`
}

type ChartData struct {
	ID     int    `json:"id"`
	XValue string `json:"x_value"`
	YValue string `json:"y_value"`
}
