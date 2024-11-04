import { useEffect, useState } from "react"

const App = () => {
  const [todos, setTodos] = useState<any>([])
  const [newTodo, setNewTodo] = useState({
    id: 1,
    task: "",
    iscompleted: false
  })

  async function postTodo(e: any) {
      e.preventDefault()
      fetch("http://localhost:3080/todos", {
        method: "POST",
        body: JSON.stringify({
          task: newTodo.task,
          iscompleted: false
        }),
        headers: {
          "Content-type": "application/json; charset=UTF-8"
        }
      }).then(resp => resp.json()).then(data => setNewTodo({
        ...newTodo,
        id: data.id
      }))

      setTodos([...todos, newTodo])
  }

  useEffect(() => {
    fetch("http://localhost:3080/todos")
      .then(resp => resp.json())
      .then(data => setTodos(data))
  }, [])

  return (
    <>
      <h1 className="text-red-500">Gotodo</h1>
      <form>
        <input type="text" onChange={(event) => setNewTodo({
          ...newTodo,
          task: event.target.value
        })} placeholder="Enter a task..." />
        <input type="submit" onClick={postTodo} />
      </form>
      <ul>
        {
         todos.map((todo: any) => (
          <li key={todo.id}>{todo.id}</li>
         ))
        }
      </ul>
    </>
  )
}

export default App
