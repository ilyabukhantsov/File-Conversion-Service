import { useFileUpload } from "@/modules/file";

export const FileUploadPage = () => {
  const {
    file,
    loading,
    message,
    selectFile,
    removeFile,
    submit,
  } = useFileUpload();

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-6 rounded-2xl shadow-md w-full max-w-md">
        <h1 className="text-xl font-semibold mb-4 text-center">
          Завантаження файлу
        </h1>

        {file ? (
          <div className="flex justify-between gap-2 items-center mb-4">
            <p className="text-sm text-gray-600 overflow-hidden">{file.name}</p>
            <button onClick={removeFile} disabled={loading}>✕</button>
          </div>
        ) : (
          <input
            type="file"
            accept=".pdf,.doc,.docx"
            onChange={(e) =>
              selectFile(e.target.files ? e.target.files[0] : null)
            }
            className="mb-4 w-full"
          />
        )}

        <button
          onClick={submit}
          disabled={loading}
          className="w-full bg-blue-600 text-white py-2 rounded-xl hover:bg-blue-700 disabled:opacity-50"
        >
          {loading ? "Завантаження..." : "Відправити"}
        </button>

        {message && (
          <p className="mt-4 text-center text-sm text-gray-700">
            {message}
          </p>
        )}
      </div>
    </div>
  );
}
