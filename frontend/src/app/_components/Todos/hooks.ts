import { graphql } from "@/lib/__graphql__";
import { useQuery } from "@tanstack/react-query";
import { request } from "graphql-request";

export const useTodosQuery = () => {
  return useQuery({
    queryKey: ["todos"],
    queryFn: async () => {
      return await request(import.meta.env.VITE_GRAPHQL_URL, GET_TODOS);
    },
  });
};

const GET_TODOS = graphql(/* GraphQL */ `
  query All_Todos {
    todos {
      id
      text
      done
      createdAt
    }
  }
`);
