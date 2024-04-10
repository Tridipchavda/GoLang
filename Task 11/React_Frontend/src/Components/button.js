export const Button = ({ onClick, children }) => (
    <button
      className="bg-blue-500 mx-10 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      onClick={onClick}
    >
      {children}
    </button>
  );
  