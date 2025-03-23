import { graphql } from "@/lib/__graphql__";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import request from "graphql-request";

export const useDeleteTodoMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: deleteTodo,
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });
};

const deleteTodo = async ({ id }: { id: string }) => {
  return await request(window.location.origin + "/query", DELETE_TODO, {
    todoId: id,
  });
};

const DELETE_TODO = graphql(`
  mutation DeleteTodo($todoId: ID!) {
    deleteTodo(todoId: $todoId) {
      id
      text
      done
      createdAt
    }
  }
`);
