import { useEffect, useState } from "react";
import type { ApiResponse, User } from "./types.ts";

const API_BASE = "http://localhost:8080";

function App() {
  const [health, setHealth] = useState<string>("");
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchData() {
      try {
        const healthRes = await fetch(`${API_BASE}/health`);
        const healthData = await healthRes.json();

        const usersRes = await fetch(`${API_BASE}/users`);
        const usersResData: ApiResponse = await usersRes.json();
        const usersData: User[] = usersResData.data;

        setHealth(healthData.status);
        setUsers(usersData);
      } catch (err) {
        console.error("Error fetching data", err);
      } finally {
        setLoading(false);
      }
    }

    fetchData();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div style={{ padding: "2rem", fontFamily: "sans-serif" }}>
      <h1>Welcome</h1>

      <section>
        <h2>Server Health</h2>
        <p>
          Status: <strong>{health}</strong>
        </p>
      </section>

      <section>
        <h2>Users</h2>
        <ul>
          {users.map((user) => (
            <li key={user.id}>{user.name}</li>
          ))}
        </ul>
      </section>
    </div>
  );
}

export default App;
