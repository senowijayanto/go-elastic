## Overview
The `elk` folder contains configuration files and Docker setup for running the ELK stack (Elasticsearch, Logstash, Kibana).

## Structure
```
├── docker-compose.yml    # Docker Compose setup for ELK
├── elasticsearch         # Elasticsearch configurations
├── kibana                # Kibana configurations
└── logstash              # Logstash configurations
```

## Running the ELK Stack
1. Navigate to the `elk` folder:
   ```bash
   cd elk
   ```
2. Start the ELK stack:
   ```bash
   docker-compose up -d
   ```
3. Access Kibana at [http://localhost:5601](http://localhost:5601).

## Notes
- Ensure Docker is installed and running.
- Update the `logstash.conf` pipeline as needed to match your application's logging format.
