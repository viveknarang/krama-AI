import React from 'react';
import Paper from '@material-ui/core/Paper'; 
import Tabs from '@material-ui/core/Tabs'; 
import Tab from '@material-ui/core/Tab'; 

class Footer extends React.Component {

  render() {
    return (
      <div className="Footer">
        <Paper square>
          <Tabs
            indicatorColor="primary"
            textColor="primary"
            aria-label="disabled tabs example"
          >
            <Tab label="Company Name " disabled/>
            <Tab label="Legal" />
            <Tab label="Return Policy" />
          </Tabs>
        </Paper>
      </div>
    )
  };

}

export default Footer;
