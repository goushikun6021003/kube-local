package model

type RecvData struct {
	Rules Ruleser `json:"rules"`
	DstObj int `json:"dst_obj"`
}

type Ruleser struct {
	Id          int64  `orm:"column(id);auto" json:"id,omitempty"`
	Expr        string `orm:"column(expr);size(1023)" json:"expr"`
	Op          string `orm:"column(op);size(31)" json:"op"`
	Value       string `orm:"column(value);size(1023)" json:"value"`
	For         string `orm:"column(for);size(1023)" json:"for"`
	Summary     string `orm:"column(summary);size(1023)" json:"summary"`
	Description string `orm:"column(description);size(1023)" json:"description"`
	Prom        *Proms `orm:"rel(fk)" json:"prom_id"`
	Plan        *Plans `orm:"rel(fk)" json:"plan_id"`
}

type Proms struct {
	Id   int64  `orm:"auto" json:"id,omitempty"`
	Name string `orm:"size(1023)" json:"name"`
	Url  string `orm:"size(1023)" json:"url"`
}

type Plans struct {
	Id          int64  `orm:"auto" json:"id,omitempty"`
	RuleLabels  string `orm:"column(rule_labels);size(255)" json:"rule_labels"`
	Description string `orm:"column(description);size(1023)" json:"description"`
}

