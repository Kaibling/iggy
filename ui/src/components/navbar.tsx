import { Link } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import './Navbar.css';

const Navbar = () => {
  const navigate = useNavigate();
  const logout = async () => {
    const apiUrl = import.meta.env.VITE_API_URL + "/auth/logout";
    const token = localStorage.getItem("rh_token");
    try {
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ' + token,
        },
      });
      if (response.ok) {
        const result = await response.json();
        if (result.success) {
          localStorage.removeItem("rh_token");
          console.log('Success:', result);
          navigate("/home");
        } else {
          console.error('Error2:', response);
        }

      } else {
        console.error('Error:', response);
        // Handle error
      }
    } catch (error) {
      console.error('Error:', error);
      // Handle error
    }
  }

  return (
    <nav className="navbar">
      <div className="navbar-logo">
        <Link to="/">iggy</Link>
      </div>
      <ul className="navbar-links">
        <li>
          <Link to="/workflows">Workflows</Link>
        </li>
        <li>
          <Link to="/workflow-backup">WF Backup</Link>
        </li>
        <li>
          <Link to="/runs">Runs</Link>
        </li>
        <li>
          <Link to="/dynamic-tables">DynTabs</Link>
        </li>
        <li>
          <Link to="/users">Users</Link>
        </li>
        <li>
          <Link to="/tokens">Tokens</Link>
        </li>
        <li>
          <a onClick={logout}>Logout</a>
        </li>
      </ul>
    </nav>
  );
};

export default Navbar;