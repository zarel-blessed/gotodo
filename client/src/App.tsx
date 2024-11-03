import { useEffect, useState } from "react"

const App = () => {
  const [todos, setTodos] = useState([])

  useEffect(() => {
    fetch("http://localhost:3080/todos")
      .then(resp => resp.json())
      .then(data => setTodos(data))
  }, [])

  return (
    <>
      <h1 className="text-red-500">Hello, FullStack</h1>
      <ul>
        {
         todos.map((todo: any) => (
          <li key={todo.id}>{todo.task}</li>
         ))
        }
      </ul>
    </>
  )
}

export default App
