#!/bin/bash

validate_project_name() {
  if [[ $1 =~ ^[a-zA-Z0-9\-]+$ && ! $1 =~ \  ]]; then
    return 0
  else
    return 1
  fi
}

create_simple_project() {
  project_name=$1
  echo "creating $project_name project..."
  mkdir -p "$project_name"
  cd "$project_name" || exit
  go mod init "$project_name"
  cat <<EOL > main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello from $project_name!")
}
EOL
  echo "project $project_name created successfully!"
}

create_api_project() {
  project_name=$1
  echo "creating $project_name project..."
  mkdir -p "$project_name"
  cd "$project_name" || exit
  go mod init "$project_name"
  mkdir -p routes models repositories services configs controllers
  cat <<EOL > main.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting API server for $project_name...")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from $project_name API!")
    })
    http.ListenAndServe(":8080", nil)
}
EOL
  echo "project $project_name created successfully with routes, models, repositories, services, configs, and controllers!"
}

echo "Go Project Generator"
read -rp "enter project name: " project_name

if ! validate_project_name "$project_name"; then
  echo "Invalid project name! Please use only letters, numbers, and hyphens (-), no spaces."
  exit 1
fi

read -rp "choose project type (simple, api): " project_type

if [[ "$project_type" == "simple" ]]; then
  create_simple_project "$project_name"
elif [[ "$project_type" == "api" ]]; then
  create_api_project "$project_name"
else
  echo "Invalid project type! Please choose 'simple' or 'api'."
  exit 1
fi
