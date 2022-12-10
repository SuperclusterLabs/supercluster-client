import { useRouteError } from "react-router-dom";

function NotFound() {
  const error = useRouteError();
  console.error(error);

  return <h1>Page not found</h1>;
}

export default NotFound;
