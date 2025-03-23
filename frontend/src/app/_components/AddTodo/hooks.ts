import { graphql } from "@/lib/__graphql__";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import request from "graphql-request";

export const useCreateTodoMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: createNewTodo,
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });
};

const createNewTodo = async ({ text }: { text: string }) => {
  return await request(import.meta.env.VITE_GRAPHQL_URL, CREATE_TODO, {
    newTodo: {
      text,
    },
  });
};

const CREATE_TODO = graphql(`
  mutation CreateTodo($newTodo: NewTodo!) {
    createTodo(newTodo: $newTodo) {
      id
      text
      done
    }
  }
`);
