import Navbar from './components/navbar';
import {  Routes, Route, Navigate } from 'react-router-dom';
import WorkflowsList from './components/workflow/workflows';
import WorkflowNew from './components/workflow/workflowNew';
import WorkflowDetail from './components/workflow/workflowDetail';
import Runlist from './components/run/runlist';
import RunDetail from './components/run/runDetail';
import './App.css'



function App() {
  // const [count, setCount] = useState(0)

  return (
    <div>
      <Navbar />
      <Routes>
        <Route path="/workflows" element={<WorkflowsList />} />
       
        <Route path="/workflows/new" element={<WorkflowNew />} />
        <Route path="/workflows/:id" element={<WorkflowDetail />} />
        <Route path="/runs" element={<Runlist />} />
        <Route path="/runs/:id" element={<RunDetail />} />
        <Route path="*" element={<Navigate to="/workflows" />} />
      </Routes>
      </div>
  )
}

export default App;
