import React from 'react';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';

class Header extends React.Component {

  render() {
    return (
      <React.Fragment>


        <Grid container xs={12} md={12} lg={12} xl={12} className="Header">


          <Grid item xs={12} md={9} lg={9} xl={9}>

            <span  className="HeaderTitle"><img alt="#" width="350px" height="80px" src="https://via.placeholder.com/350x80"/></span>

          </Grid>

          <Grid item xs={12} md={3} lg={3} xl={3}>

            <Button variant="contained" startIcon={<ShoppingCartIcon />} color="primary" className="HeaderCheckoutButton">Shopping Cart</Button>

          </Grid>

        </Grid>

        </React.Fragment>
    )
  };

}

export default Header;
