import { useUpdateTodoMutation } from "@/app/_components/Todos/Actions/TodoCheckbox/hooks";
import { Todo } from "@/lib/__graphql__/graphql";
import { toast } from "sonner";

export default function TodoCheckbox({ todo }: { todo: Todo }) {
  const { mutate, status } = useUpdateTodoMutation();

  const onSubmit = () => {
    mutate(
      {
        todo: {
          ...todo,
          done: !todo.done,
        },
      },
      {
        onSuccess: (data) => {
          toast(`You have successfully ${data.updateTodo.text}!`);
        },

        onError: () => {
          toast.error("Sorry, something went wrong.");
        },
      },
    );
  };

  return (
    <input
      type="checkbox"
      checked={todo.done}
      onClick={onSubmit}
      disabled={status === "pending"}
    />
  );
}
