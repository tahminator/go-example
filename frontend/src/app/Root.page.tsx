import TodoAddForm from "@/app/_components/AddTodo/AddTodo";
import Todos from "@/app/_components/Todos/Todos";

export default function RootPage() {
  return (
    <div className="flex flex-col items-center justify-center p-4">
      <span>Todo App with Golang {"&"} React</span>
      <TodoAddForm />
      <Todos />
    </div>
  );
}
