import React from "react";
import {MainLayout} from "./layouts/MainLayout";
import {TasksState} from "./components/context/tasks/TasksState";
import {TaskModalState} from "./components/context/taskModal/TaskModalState";


function App() {
    return (
        <TaskModalState>
            <TasksState>
                <div className="App">
                    <MainLayout/>
                </div>
            </TasksState>
        </TaskModalState>

    );
}

export default App;
