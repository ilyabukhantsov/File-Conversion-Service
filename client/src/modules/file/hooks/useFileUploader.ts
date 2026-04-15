import { useState } from "react";
import { validateFile } from "../utils/validateFile";
import { uploadFile } from "../api/fileUpload";

export const useFileUpload = () => {
  const [file, setFile] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState<string>("");

  const selectFile = (file: File | null) => {
    setFile(file);
    setMessage("");
  };

  const removeFile = () => {
    setFile(null);
  };

  const submit = async () => {
    const error = validateFile(file);

    if (error) {
      setMessage(error);
      return;
    }

    try {
      setLoading(true);
      setMessage("");

      await uploadFile(file!);

      setMessage("Файл успішно завантажено");
    } catch (e: any) {
      setMessage(e.message);
    } finally {
      setLoading(false);
    }
  };

  return {
    file,
    loading,
    message,
    selectFile,
    removeFile,
    submit,
  };
};