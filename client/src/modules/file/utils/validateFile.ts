export const validateFile = (file: File | null) => {
  if (!file) return "Будь ласка, оберіть файл";

  const allowedTypes = [
    "application/pdf",
    "application/msword",
  ];

  if (!allowedTypes.includes(file.type)) {
    return "Непідтримуваний формат файлу";
  }

  return null;
};