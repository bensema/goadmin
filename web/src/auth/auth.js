import { Navigate } from "react-router";

export function RequireAuth({ children }) {
  const authed = localStorage.getItem("login");

  return authed === "true" ? children : <Navigate to="/sign-in" replace />;
}
