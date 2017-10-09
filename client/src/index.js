import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter, Switch, Route } from 'react-router-dom';

import Home from './App';
import Add from './Add';

ReactDOM.render(
  <BrowserRouter>
    <Switch>
      <Route exact path="/" component={Home} />
      <Route exact path="/add" component={Add} />
    </Switch>
  </BrowserRouter>
  , document.getElementById('app'));
