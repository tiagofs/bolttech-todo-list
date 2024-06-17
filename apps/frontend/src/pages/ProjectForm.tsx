import { useState } from "react"
import { createProject } from "../lib/backend-api"

export default function ProjectForm () {
  const [projectName, setProjectName] = useState("")

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const result = await createProject(projectName);
    console.log('result: ', result)
  }

  return (
    <>
    <h1>Create New Project</h1>
    <form onSubmit={handleSubmit}>
      <input
        value={projectName}
        onChange={(e) => setProjectName(e.target.value)}
        placeholder="Project Name"
      ></input>
      <button>Create Project</button>
    </form>
    </>
  )
}