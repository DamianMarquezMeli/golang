import React from 'react'
import { Redirect, Route, Switch } from 'react-router'
import { Adminpanel } from '../adminpanel/Adminpanel'
import { Edit } from '../edit/Edit'
import { Map } from '../map/Map'
import { Navbar } from '../navbar/Navbar'
import { Register } from '../register/Register'

export const DashboardRoutes = () => {
    return (
        <>

            <Navbar/>
            <div>
                <Switch>
                    <Route exact path='/register' component={Register}/>
                    <Route exact path='/' component={Map}/>
                    <Route exact path='/adminpanel' component={Adminpanel}/>
                    <Route exact path='/edit' component={Edit}/>
                    <Redirect to='/'/>
                </Switch>
            </div>
          
            
        </>
    )
}
