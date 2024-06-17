import { useContext } from "react";
import Navbar from "./components/navbar";
import { AuthContext } from "./utils/AuthContext";
import TodosList from "./pages/TodosList";
import Login from "./pages/Login";
import ProjectForm from "./pages/ProjectForm";

export default function App() {
  const authContext = useContext(AuthContext);

  if (!authContext) {
    throw new Error('AuthContext must be used within an AuthProvider');
  }

  const { authState } = authContext;

  
  return (
      <div className="flex flex-col">
        <Navbar />
        <main className="min-w-screen w-full mt-[50px]">
          {/* {authState ? <TodosList /> : <Login />} */}
          <ProjectForm />
        </main>
      </div>
  );
}
