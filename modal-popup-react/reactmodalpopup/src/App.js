import React, { Fragment } from 'react';  
import ModalPopup from './src/modal_popup';  
  
class App extends React.Component {  
  constructor() {  
    super();  
    this.state = {  
      showModalPopup: false  
    }  
  }  
  isShowPopup = (status) => {  
    this.setState({ showModalPopup: status });  
  };  
  render() {  
    return (  
      <Fragment>  
        <h3 align="center">Demo of Modal Pop up in Reactjs</h3>  
        <header align="center">  
          <Fragment>  
            <div  
              className="nav-item"  
              onClick={() => this.isShowPopup(true)}>  
              <button>Modal Pop up</button>  
            </div>  
          </Fragment>  
        </header>  
        <ModalPopup  
          showModalPopup={this.state.showModalPopup}  
          onPopupClose={this.isShowPopup}  
        ></ModalPopup>  
      </Fragment>  
    )  
  }  
}  
  
export default App;