import { useState } from "react";
import { login } from "../lib/backend-api";
import { useCookies } from "react-cookie";
import { sessionOptions } from "../utils/config";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState<string | null>(null); // New state for handling error messages
  const [cookies, setCookies] = useCookies([sessionOptions.cookieName]);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await login(email, password);
      console.log("response: ", response);
      setCookies(sessionOptions.cookieName, response);
    } catch (error) {
      console.error("Authentication failed:", error);
      setErrorMessage("An unexpected error occurred. Please try again.");
    }
  };

  return (
    <div className="flex flex-col max-w-lg m-auto">
      <h2>Login</h2>
      <div className="max-w-lg">
        {errorMessage && <div style={{ color: "red" }}>{errorMessage}</div>}{" "}
        <form className="flex flex-col" onSubmit={handleSubmit}>
          <input
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Email"
            className="border-2 border-slate-700 rounded py-[4px] my-[4px]"
          />
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password"
            className="border-2 border-slate-700 rounded py-[4px] my-[4px]"
          />
          <button type="submit">Login</button>
        </form>
      </div>
      <p>Don't have an account yet? Click here to register</p>
    </div>
  );
}
