OpenSearch Experiment
====================

# Installation Guidelines

- Run `docker-compose.yaml` in local by run command `docker-compose up -d`

# Getting Start
Elasticsearch has total 4 important modules

1. Cluster
2. Node
3. Index
4. Shard

## Cluster
**Cluster** is the group of nodes

## Node
**Node** is the VM that installed the elasticsearch service

## Index
**Index** is the configuration how we keep the data

## Shard
**Shard** is the module that store the data in the node

# Create Index

To create index we need to define the index structure for elasticsearch in format like this

```json
{
  "settings": {
    "index": {
      "number_of_shards": 1,
      "number_of_replicas": 1
    }
  },
  "mappings": {
    "properties": {
      "field1": {
        "type": "text|keyword|integer|double...",
        "analyzer": "english|thai"
      }
    }
  }
}
```

### Setting

- **number_of_shards** is the number of the primary shard that elasticsearch will generate to store the data
  - Cannot be edited after generated
- **number_of_replicas** is the number of the replicas shard that elasticsearch will generate to be the replicas of the data in other nodes
  - The number of replicas shard will be related to the primary shard
  - example: Primary Shard = 3 and number_of_replicas = 2 -> we will get total 6 Replicas Shard
  - Can be edited after generated

### Mapping

- **properties** is the structure of the data like schema in the RDBMS
  - we can also choose the `text analyzer` for the data

## Search

```http request
GET http://ELASTICSEARCH_HOST/INDEX_NAME/_search?q=ATTRIBUTE_NAME:YOUR_QUERY
```

### Example

```http request
GET http://localhost:9200/course/_search?q=courseNameEn:INDEPENDENT
```
