import React from 'react';
import './App.css';
import Header from './components/Header'
import Navigator from './components/Navigator'
import Product from './components/Product'
import Footer from './components/Footer'
import { BrowserRouter, Route, Switch } from 'react-router-dom';


const key = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjeHMiOiJmZmFhYmJjY2RkIiwiZXhwIjoxNTgwNzQ4NjAwLCJpYXQiOjE1ODA2Njg2MDAsIm5iZiI6MTU4MDY2ODUwMCwidWlkIjoiMmE0NDY0NWItYzdjOS00ZTliLTgyNmYtMWZjMjQwOTgyMjk3In0.Hn7ei2TwWKn7vdDGHE7U_CEN8KTMcGueL9lrWIiyz5A"

class App extends React.Component {

  state = {
  }

  componentDidMount() {
  }

  render() {
    return (
      <BrowserRouter>
        <div className="App">
          <Header />
          <Navigator />

              <Switch>
                <Route path='/products/:PGID' render={(props) => (<Product {...props} APIKey={key} />)} />
              </Switch>

          <Footer />
        </div>
      </BrowserRouter>
    )
  };

}

export default App;
