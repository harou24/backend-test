#!/bin/bash

# Set the base URL for the API
BASE_URL=http://localhost:50010

# Function to test API GET endpoint
test_get() {
  local endpoint=$1
  local description=$2

  response=$(curl -s -o /dev/null -w "%{http_code}" -X GET "${BASE_URL}${endpoint}" -H 'Accept: application/json')
  if [ "$response" == "200" ]; then
    echo "Test ${description}: PASS"
  else
    echo "Test ${description}: FAIL"
    echo "Expected: 200"
    echo "Actual:   $response"
  fi
}

# Function to test API POST endpoint
test_post() {
  local endpoint=$1
  local data=$2
  local description=$3

  response=$(curl -s -o /dev/null -w "%{http_code}" -X POST "${BASE_URL}${endpoint}" -H 'Content-Type: application/json' -d "${data}")
  if [ "$response" == "201" ]; then
    echo "Test ${description}: PASS"
  else
    echo "Test ${description}: FAIL"
    echo "Expected: 201"
    echo "Actual:   $response"
  fi
}

# Function to test API PUT endpoint
test_put() {
  local endpoint=$1
  local data=$2
  local description=$3

  response=$(curl -s -o /dev/null -w "%{http_code}" -X PUT "${BASE_URL}${endpoint}" -H 'Content-Type: application/json' -d "${data}")
  if [ "$response" == "200" ]; then
    echo "Test ${description}: PASS"
  else
    echo "Test ${description}: FAIL"
    echo "Expected: 200"
    echo "Actual:   $response"
  fi
}

# Function to test API DELETE endpoint
test_delete() {
  local endpoint=$1
  local description=$2

  response=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "${BASE_URL}${endpoint}")
  if [ "$response" == "200" ]; then
    echo "Test ${description}: PASS"
  else
    echo "Test ${description}: FAIL"
    echo "Expected: 200"
    echo "Actual:   $response"
  fi
}

# Test cases for reading
test_get "/search?species=cat" "searching by species"
test_get "/search?weight=5000" "searching by weight"
test_get "/search?species=cat&weight=5000" "searching by species and weight"

# Test cases for creating
new_breed='{"species":"dog","pet_size":"medium","name":"Beagle","average_male_adult_weight":10000,"average_female_adult_weight":9000}'
test_post "/breeds" "${new_breed}" "creating a new breed"

# Test cases for updating
update_breed='{"species":"dog","pet_size":"medium","name":"Beagle","average_male_adult_weight":10500,"average_female_adult_weight":9500}'
test_put "/breeds/1" "${update_breed}" "updating a breed with ID 1"

# Test cases for deleting
test_delete "/breeds/1" "deleting a breed with ID 1"

