import { API_URL } from "@/shared/api";

export const uploadFile = async (file: File) => {
  const formData = new FormData();
  formData.append("file", file);

  const response = await fetch(`${API_URL}/files/upload`, {
    method: "POST",
    body: formData,
  });

  if (!response.ok) {
    throw new Error("Помилка при завантаженні");
  }

  return response.json();
};