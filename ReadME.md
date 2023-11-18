
# <p align="center"> Golang Standard Rest-API </p>

    This project is to include every basics component that usually included in backend Golang project.

## Table of Contents
- [ Golang Standard Rest-API ](#-golang-standard-rest-api-)
  - [Table of Contents](#table-of-contents)
  - [Project Component](#project-component)
  - [Project Structure](#project-structure)
    - [1. `cmd`](#1-cmd)
    - [2. `handler`](#2-handler)
    - [3. `db`](#3-db)
    - [4. `service`](#4-service)
    - [5. `Model`](#5-model)
  - [Tutorial to create new module/package](#tutorial-to-create-new-modulepackage)

## Project Component

    - Modular loosely coupled design
    - Rest API back end service (GIN Framerwork)
    - Event Driven Architecture (Kafka)
    - SQL Database as main datastore (Postgresql)
    - NoSQL Datase as cache and search (Redis,Elasticsearch)
    - CI (Github action)
    - CD (AWS EKS)
    - Project management (JIRA)
    - Logging (Logrus & EFK stack & Docker driver)
    - Database Migration (golang-migrate with rollback mechanism)
    - Application Configuration (JSON file + Env Variable)
    - Validator (Go Playground)
    - Project Repository (Github)
    - Docker Image Registry (Docker Hub)
    - High Availability (Database replication + EKS)
    - Scheduler (gocron)
    - External Integration
    - API documentation (Postman)
    - Health Checker API
    - Automated Testing
      - Jmeter (Load + stress test)
      - Functional Testing (Postman)
      - Regression Testing
      - Testing management tools
      - Unit Testing (Testify)

## Project Structure

    This part will outlines the structure of project providing an overview of the purpose of each directory.

### 1. `cmd`
    Logic to Run Your application and load application config

    **Scope & Do's** 
    1. Init Scheduler
    2. Init Route
    3. Init Database connection
    4. Init logging 

    **Dont's**
    1. Hardcode anything
### 2. `handler`
    dilarang ada bisnis logic here , response building, request mapping, 1st layer validator
### 3. `db`
    all related retrieving data from database handled here.
### 4. `service`
    put your bussiness logic here , 2nd layer validator
### 5. `Model`
    define DTO, DTIN, database model. Create custom method from each model to do specific job such as validation



## Tutorial to create new module/package
For example you want to create a basic CRUD feature for object called "Course"

1. Define Course Model. Make a struct out of it and place it on "model" directory. create a new file called course.go
2. inside handler directory. Copy template.txt and paste it as coursehandler.go
3. Edit and modify it until there is no error
4. Go to test package. Copy testtemplate folder and rename it.
5. rename 