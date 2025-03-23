import { z } from "zod";

export const createTodoSchema = z.object({
  text: z
    .string()
    .trim()
    .min(1, { message: "The todo must be atleast 1 character long." })
    .max(150, {
      message: "The todo cannot be longer than 150 characters long.",
    }),
});
