import React from 'react';
import Paper from '@material-ui/core/Paper'; 
import Tabs from '@material-ui/core/Tabs'; 
import Tab from '@material-ui/core/Tab'; 


class Navigator extends React.Component {

  render() {
    return (
      <div className="Navigator">

        <Paper square>
          <Tabs
            indicatorColor="primary"
            textColor="primary"
            aria-label="disabled tabs example"
          >
            <Tab label="Departments" />
            <Tab label="Deals & Offers"/>
            <Tab label="Trending Products" />
          </Tabs>
        </Paper>

      </div>
    )
  };

}

export default Navigator;
