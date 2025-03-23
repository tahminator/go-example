import { useDeleteTodoMutation } from "@/app/_components/Todos/Actions/DeleteTodo/hooks";
import { toast } from "sonner";

export default function TodoDeleteButton({ id }: { id: string }) {
  const { mutate, status } = useDeleteTodoMutation();

  const onSubmit = () => {
    const toastId = toast.loading("Please wait, deleting todo...");
    mutate(
      { id },
      {
        onSuccess: (data) => {
          toast.success(
            `Todo with ID of ${data.deleteTodo.id} has been deleted.`,
            {
              id: toastId,
            },
          );
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
