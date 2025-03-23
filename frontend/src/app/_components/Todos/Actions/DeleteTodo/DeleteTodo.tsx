import { useDeleteTodoMutation } from "@/app/_components/Todos/Actions/DeleteTodo/hooks";
import { toast } from "sonner";

export default function TodoDeleteButton({ id }: { id: string }) {
  const { mutate, status } = useDeleteTodoMutation();

  const onSubmit = () => {
    const toastId = toast.loading("Please wait, deleting todo...");
    mutate(
      { id },
      {
        onSuccess: ({ deleteTodo: { text } }) => {
          toast.success(`"${text}" has been removed from the list.`, {
            id: toastId,
          });
        },
      },
    );
  };

  return (
    <button disabled={status === "pending"} onClick={onSubmit}>
      🗑️
    </button>
  );
}
