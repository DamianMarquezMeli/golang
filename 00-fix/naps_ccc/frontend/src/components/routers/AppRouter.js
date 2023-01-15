import React from 'react'
import { Route, BrowserRouter as Router, Switch, Redirect } from 'react-router-dom'
// import { Login } from '../login/Login'
// import { DashboardRoutes } from './DashboardRoutes'
import { Map } from '../map/Map'

export const AppRouter = () => {
    return (
        <Router>
            <div>
                <Switch>
                    <Route exact path='/' component={Map}/>
                    <Redirect to='/'/>
                    {/* <Route exact path='/login' component={Login}/> */}
                    {/* <Route path='/' component={DashboardRoutes}/> */}
                </Switch>
            </div>
        </Router>
    )
}
