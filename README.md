# Syndio Senior BE Takehome Assignment - Gayle Angeles

Provides a simple API to manage employee information.

## Prerequisites

- [Docker](https://www.docker.com/get-started) must be installed and running on your system.

## Getting Started

### 1. Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/gayleangeles/senior-backend-engineer-exercise.git
cd senior-backend-engineer-exercise
```

### 2. Build the Docker Image

In the project directory, run the following command to build the Docker image:

```powershell
docker build -t angeles .
```

### 2. Run the Docker Image

After the image is built, start the container using:

```powershell
docker run -e PORT=8080 -p 8080:8080 angeles
```

The application will be available at http://localhost:8080 if it starts successfully.

# Running the container with a custom port

To change the exposed port, set the `PORT` environment variable when running the container, `8083` below being the specified port:

```powershell
docker run -e PORT=8083 -p 8080:8083 angeles
```

# API Documentation

### 1. Get Employees

Retrieve the entire list of employees.

- **Endpoint**: `GET /employees`

#### Response Schema

- **Type**: Array of `Employee` objects
- **`Employee` object**:
  - **id** (integer): Unique identifier for the employee
  - **gender** (string): Gender of the employee

#### Example Response

```json
[
  {
    "id": 1,
    "gender": "male"
  },
  {
    "id": 2,
    "gender": "male"
  },
  {
    "id": 3,
    "gender": "male"
  },
  {
    "id": 4,
    "gender": "female"
  },
  {
    "id": 5,
    "gender": "female"
  },
  {
    "id": 6,
    "gender": "female"
  },
  {
    "id": 7,
    "gender": "non-binary"
  }
]
```

---

### 2. Add a Job to an Employee

Assign a job to a specific employee by their `id`.

- **Endpoint**: `POST /employees/{id}/jobs`
- **Path Parameter**:
  - `id` _(integer)_: Unique identifier for the employee

#### Request Body Schema

- **Properties**:
  - **department** (string): The department where the employee will work
  - **job_title** (string): The job title assigned to the employee

### Example Request Body

```json
{
  "department": "Engineering",
  "jobTitle": "Senior Enginer"
}
```

#### Response Schema

- **Type**: Object
- **Properties**:
  - **id** _(integer)_: The ID of the newly added job
  - **message** _(string)_: Confirmation message for successful job assignment

#### Example Response

```json
{
  "id": 1,
  "message": "employee job added successfully"
}
```

---

### 3. Update an Employee's Job

Update the job information for a specific employee's job by job `id`.

- **Endpoint**: `PATCH /employees/jobs/{id}`
- **Path Parameter**:

  - `id` _(integer)_: The job ID of the employee's job to update.

  #### Request Body Schema

  - **Type**: Object
  - **Properties**:
    - **department** _(string)_: The new department for the job
    - **job_title** _(string)_: The new job title

  #### Example Request Body

  ```json
  {
    "department": "My New Updated Department",
    "job_title": "My New Updated Job Title"
  }
  ```

- **Response Body**:

  #### Response Schema

  - **Type**: Boolean
  - **Description**: `true` if the job was successfully updated, otherwise an error response

  #### Example Response Body

  ```json
  true
  ```

---

### 4. Get All Employee Jobs

Retrieve a list of all jobs assigned to employees.

- **Endpoint**: `GET /employees/jobs`

  #### Response Schema

  - **Type**: Array of `EmployeeJob` objects
  - **`EmployeeJob` object**:
    - **id** _(integer)_: Unique identifier for the job
    - **employee_id** _(integer)_: ID of the employee to whom the job is assigned
    - **department** _(string)_: The department where the job is assigned
    - **job_title** _(string)_: The title of the assigned job

  #### Example Response

  ```json
  [
    {
      "id": 1,
      "employee_id": 2,
      "department": "My Department",
      "job_title": "My Job Title"
    },
    {
      "id": 2,
      "employee_id": 2,
      "department": "Updated Department",
      "job_title": "Updated Job Title"
    }
  ]
  ```

---

### 5. Get Jobs for a Specific Employee

Retrieve a list of jobs assigned to a specific employee by their `id`.

- **Endpoint**: `GET /employees/{id}/jobs`
- **Path Parameter**:
  - `id` _(integer)_: The employee's ID whose jobs are to be retrieved.

#### Response Schema

- **Type**: Array of `EmployeeJob` objects
- **`EmployeeJob` object**:
  - **id** _(integer)_: Unique identifier for the job
  - **employee_id** _(integer)_: ID of the employee to whom the job is assigned
  - **department** _(string)_: The department where the job is assigned
  - **job_title** _(string)_: The title of the assigned job

#### Example Response

```json
[
  {
    "id": 1,
    "employee_id": 2,
    "department": "My Department",
    "job_title": "My Job Title"
  },
  {
    "id": 2,
    "employee_id": 2,
    "department": "Updated Department",
    "job_title": "Updated Job Title"
  }
]
```
