package model

type RecvData struct {
	Rules  Ruleser `json:"rules"`
	DstObj string  `json:"dst_obj"`
}

type Ruleser struct {
	Id          int64  `json:"id,omitempty"`
	Expr        string `json:"expr"`
	Op          string `json:"op"`
	Value       string `json:"value"`
	For         string `json:"for"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Prom        *Proms `json:"prom_id"`
	Plan        *Plans `json:"plan_id"`
}

type Proms struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Plans struct {
	Id          int64  `json:"id,omitempty"`
	RuleLabels  string `json:"rule_labels"`
	Description string `json:"description"`
}
