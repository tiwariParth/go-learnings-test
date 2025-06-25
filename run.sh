#!/bin/bash

# Helper script for running Go projects in the repository

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Check if a week number is provided
if [ $# -lt 1 ]; then
    echo -e "${RED}Error: Please specify a week number${NC}"
    echo "Usage: $0 <week-number> [project-name]"
    echo "Example: $0 1 todo"
    echo "Example: $0 4-5 url-checker"
    exit 1
fi

# Week directory
WEEK_DIR="week-$1"

# Check if week directory exists
if [ ! -d "$WEEK_DIR" ]; then
    echo -e "${RED}Error: Week directory '$WEEK_DIR' does not exist${NC}"
    exit 1
fi

# If project name is provided
if [ $# -ge 2 ]; then
    PROJECT_NAME=$2
    PROJECT_DIR="$WEEK_DIR/project/$PROJECT_NAME"
    
    # Check if project directory exists
    if [ ! -d "$PROJECT_DIR" ]; then
        echo -e "${RED}Error: Project directory '$PROJECT_DIR' does not exist${NC}"
        exit 1
    fi
    
    # Navigate to project directory
    cd "$PROJECT_DIR" || exit 1
    
    echo -e "${BLUE}Running project '$PROJECT_NAME' from week $1...${NC}"
    
    # Check if there's a go.mod file
    if [ -f "go.mod" ]; then
        # Run the project
        if [ -f "main.go" ]; then
            echo -e "${GREEN}Running: go run main.go${NC}"
            go run main.go
        else
            echo -e "${RED}Error: main.go not found in project directory${NC}"
            echo "Available Go files:"
            ls -1 *.go 2>/dev/null || echo "No Go files found"
            exit 1
        fi
    else
        echo -e "${RED}Error: go.mod not found. Attempting to initialize module...${NC}"
        go mod init "github.com/parth/go-learnings/$WEEK_DIR/$PROJECT_NAME"
        go mod tidy
        
        if [ -f "main.go" ]; then
            echo -e "${GREEN}Running: go run main.go${NC}"
            go run main.go
        else
            echo -e "${RED}Error: main.go not found in project directory${NC}"
            echo "Available Go files:"
            ls -1 *.go 2>/dev/null || echo "No Go files found"
            exit 1
        fi
    fi
else
    echo -e "${BLUE}Available projects in week $1:${NC}"
    ls -1 "$WEEK_DIR/project/" 2>/dev/null || echo "No projects found"
    
    echo -e "\n${BLUE}Available exercises in week $1:${NC}"
    ls -1 "$WEEK_DIR/exercises/" 2>/dev/null || echo "No exercises found"
    
    echo -e "\n${BLUE}Available notes in week $1:${NC}"
    ls -1 "$WEEK_DIR/notes/" 2>/dev/null || echo "No notes found"
    
    echo -e "\nTo run a project: $0 $1 <project-name>"
fi
