package models

type EntityInstance struct {
    LabelName   string `json:"labelName"`
    Description string `json:"description"`
}

type Relationship struct {
    StartNode  Node                   `json:"startNode"`
    EndNode    Node                   `json:"endNode"`
    Type       string                 `json:"type"`
    Properties map[string]interface{} `json:"properties"`
}

type Node struct {
    Label      string                 `json:"label"`
    Properties map[string]interface{} `json:"properties"`
}

type SchemaRequest struct {
    Ontology      bool             `json:"ontology"`
    Data          []EntityInstance `json:"data"`
    Relationships []Relationship   `json:"relationships"`
}

type SchemaResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    struct {
        SchemaID       string `json:"schemaId"`
        EntitiesCount  int    `json:"entitiesCount"`
        RelationsCount int    `json:"relationsCount"`
    } `json:"data"`
}

type DemoResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

type StaticDataResponse struct {
    Entities      []EntityInstance `json:"entities"`
    Relationships []Relationship   `json:"relationships"`
    Stats         SystemStats      `json:"stats"`
}

type SystemStats struct {
    TotalEntities     int `json:"totalEntities"`
    TotalRelations    int `json:"totalRelations"`
    ProcessCount      int `json:"processCount"`
    BusinessRuleCount int `json:"businessRuleCount"`
}
