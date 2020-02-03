import React from 'react';

import CircularProgress from '@material-ui/core/CircularProgress';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
import Rating from '@material-ui/lab/Rating';
import Alert from '@material-ui/lab/Alert';
import Typography from '@material-ui/core/Typography';
import "react-responsive-carousel/lib/styles/carousel.min.css";
import { Carousel } from 'react-responsive-carousel';
import GradeIcon from '@material-ui/icons/Grade';
import FormControl from '@material-ui/core/FormControl'
import InputLabel from '@material-ui/core/InputLabel'
import Select from '@material-ui/core/Select'
import AddShoppingCartIcon from '@material-ui/icons/AddShoppingCart';
import LinearProgress from '@material-ui/core/LinearProgress';
import TextField from '@material-ui/core/TextField';
import Table from '@material-ui/core/Table';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TableCell from '@material-ui/core/TableCell';
import TableBody from '@material-ui/core/TableBody';
import Paper from '@material-ui/core/Paper';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import axios from 'axios';

import '../App.css';
import '../Slider.css'

let pgrequestURL = "http://127.0.0.1:9005/catalog/v1/productgroups/"
let pirequestBaseURL = "http://127.0.0.1:9005/catalog/v1/inventory/products/"

let filter = {};

class Product extends React.Component {

    state = {
        code: 0,
        productGroup: {
            Images: [""]
        },
        selectedProduct: {},
        selectedProductSKU: "",
        selectors: {},
        products: [],
        selectedProductAvailableQuantity: 0,
        loading: true,
        disabled: true,
        pageLoading: true
    }

    componentDidMount() {

        axios.get(pgrequestURL + this.props.match.params.PGID, { headers: { "x-requested-with": "", "x-access-token": this.props.APIKey } })
            .then(res => {

                if (res.data == null) {
                    return
                }

                this.setState({ productGroup: res.data.Response, code: res.data.Code, selectedProduct: res.data.Response.Products[res.data.Response.MainProductSKU] })
                this.setState({ selectors: res.data.Response.Selectors })
                this.setState({ selectedProductSKU: res.data.Response.MainProductSKU })

                axios.get(pirequestBaseURL + res.data.Response.MainProductSKU, { headers: { "x-requested-with": "", "x-access-token": this.props.APIKey } })
                    .then(res => {
                        this.setState({ selectedProductAvailableQuantity: res.data.Response.Quantity, loading: false })

                        var prds = [];
                        Object.entries(this.state.productGroup.Products).forEach(([key, val]) => {
                            prds.push(val)
                        });
                        this.setState({ products: prds })

                        Object.entries(this.state.selectedProduct.Selectors).forEach(([key, val]) => {
                            filter[key] = val
                        })

                        if (this.state.selectedProduct != null) {
                            this.setState({ disabled: false })
                            this.setState({ pageLoading: false })
                            document.title = this.state.selectedProduct.Name;
                        }

                    })

            })

    }

    loadSelectors() {

        const handleChange = event => {

            this.setState({ disabled: true })
            filter[event.target.name] = event.target.value

            var product = this.state.products.filter(item => {

                for (let key in filter) {
                    if (item.Selectors[key] === undefined || item.Selectors[key] !== filter[key])
                        return false;
                }
                return true;

            });
            if (product == null || product[0] == null) {

                this.setState({ disabled: true, selectedProductAvailableQuantity: -1 })
                this.setState({ selectedProductSKU: "-" })
                return

            } else {

                this.setState({ disabled: false })
                this.setState({ selectedProduct: product[0] })
                this.setState({ selectedProductSKU: product[0].Sku })
                this.setState({ loading: true })

                axios.get(pirequestBaseURL + product[0].Sku, { headers: { "x-requested-with": "", "x-access-token": this.props.APIKey } })
                    .then(res => { this.setState({ selectedProductAvailableQuantity: res.data.Response.Quantity, loading: false }) })

            }

        };

        var selectors = [];

        Object.entries(this.state.selectors).forEach(([key, val]) => {

            if (val.length > 1) {

                var values = [];
                for (var i = 0; i < val.length; i++) {

                    if (filter[key] === val[i]) {
                        values.push(<option key={i} value={val[i]} selected>{val[i]}</option>);
                    } else {
                        values.push(<option key={i} value={val[i]}>{val[i]}</option>);
                    }

                }

                selectors.push(
                    <FormControl key={key} className="Selector">
                        <InputLabel>{key}</InputLabel>
                        <Select native name={key} onChange={handleChange}>
                            {values}
                        </Select>
                    </FormControl>
                );

            }

        });

        return selectors;

    }

    loadProductFeatures() {

        var features = [];

        Object.entries(this.state.selectedProduct.Selectors).forEach(([key, val]) => {

            features.push(
                <TableRow key={key}>
                    <TableCell component="th" scope="row">{key}</TableCell>
                    <TableCell align="right">{val}</TableCell>
                </TableRow>
            );

        });

        Object.entries(this.state.selectedProduct.Attributes).forEach(([key, val]) => {

            features.push(
                <TableRow key={key}>
                    <TableCell component="th" scope="row">{key}</TableCell>
                    <TableCell align="right">{val}</TableCell>
                </TableRow>
            );

        });

        return features;

    }

    loadImages() {

        if (this.state.selectedProduct == null || this.state.selectedProduct.Images == null) {
            return
        }
        var images = [];
        for (var i = 0; i < this.state.selectedProduct.Images.length; i++) {
            images.push(<div key={i}><img alt="#" src={this.state.selectedProduct.Images[i]} /></div>);
        }
        return images;

    }

    render() {

        return (
            <React.Fragment>

                {this.state.pageLoading ? <div className="Centered"><br /><LinearProgress /><br />Loading Product ...</div> :

                    <Grid container xs={12} md={12} lg={12} xl={12} className="BodyGrid">

                        <Grid item xs={12} md={6} lg={6} xl={6} className="ProductImage">
                            <Carousel emulateTouch infiniteLoop={true} showThumbs={false} showArrows={true} dynamicHeight={true} >
                                {this.loadImages()}
                            </Carousel>
                        </Grid>

                        <Grid item xs={12} md={6} lg={6} xl={6} className="ProductDetails">

                            <Grid item xs={12} md={12} lg={12} xl={12}>
                                <h2>{this.state.selectedProduct.Name}</h2>
                            </Grid>

                            <Grid container xs={12} md={12} lg={12} xl={12}>
                                <Grid item xs={12} md={7} lg={7} xl={7}>
                                    <h2>{this.state.selectedProduct.Currency}&nbsp;<span className="PriceDiscount">${this.state.selectedProduct.PromotionPrice < this.state.selectedProduct.RegularPrice ? this.state.selectedProduct.PromotionPrice : this.state.selectedProduct.RegularPrice}</span></h2>
                                </Grid>
                                <Grid item xs={12} md={5} lg={5} xl={5}>
                                    {this.state.productGroup.CumulativeReviewCount !== 0 ? <Rating className="Rating" name="read-only" value={this.state.productGroup.CumulativeReviewStars || 0.0} readOnly /> : <div className="Rating" >No Reviews Yet</div>}
                                </Grid>
                            </Grid>

                            <br />

                            <Grid item xs={12} md={12} lg={12} xl={12}>
                                {this.state.loading === true ? <CircularProgress /> : null}
                                {this.state.loading === false && this.state.selectedProductAvailableQuantity === -1 ? <Alert severity="warning">Feature combination not available. Please change your feature selection...</Alert> : null}
                                {this.state.loading === false && this.state.selectedProductAvailableQuantity > 10 ? <Alert severity="success"><b>In Stock</b></Alert> : null}
                                {this.state.loading === false && this.state.selectedProductAvailableQuantity === 0 ? <Alert severity="error"><b>Out of Stock</b></Alert> : null}
                                {this.state.loading === false && this.state.selectedProductAvailableQuantity < 10 && this.state.selectedProductAvailableQuantity > 0 ? <Alert severity="warning"><b>Few left - Hurry!</b></Alert> : null}
                            </Grid>

                            <br />

                            <Grid item xs={12} md={12} lg={12} xl={12}>
                                <span className="ProductDescription">{this.state.selectedProduct.Description}</span>
                            </Grid>

                            <br /><br />

                            <Grid container xs={12} md={12} lg={12} xl={12}>
                                {this.loadSelectors()}
                            </Grid>

                            <br /><br />

                            <Grid container xs={12} md={12} lg={12} xl={12}>

                                <Grid item xs={12} md={3} lg={3} xl={3}>
                                    <TextField id="outlined-number" label="Quantity" type="number" defaultValue="1" name="Quantity" variant="outlined" className="QuantitySelector"
                                        InputLabelProps={{
                                            shrink: true,
                                        }}
                                    />
                                </Grid>

                                <Grid item xs={12} md={5} lg={5} xl={5} className="AddToBagButton">
                                    {this.state.disabled === false && this.state.selectedProductAvailableQuantity > 0 ? <Button variant="contained" color="primary" startIcon={<AddShoppingCartIcon />} fullWidth={true}>Add to Bag</Button> : <Button variant="contained" color="primary" fullWidth={true} disabled>Add to Bag</Button>}
                                </Grid>

                                <Grid item xs={12} md={4} lg={4} xl={4} className="AddToWishListButton">
                                    {this.state.disabled === false ? <Button variant="contained" color="secondary" startIcon={<GradeIcon />}>Add to Wish List</Button> : <Button variant="contained" color="secondary" startIcon={<GradeIcon />} disabled>Add to Wish List</Button>}
                                </Grid>

                            </Grid>

                            <br /><br />

                            <Grid item xs={12} md={12} lg={12} xl={12}>
                                {this.state.disabled === false ?
                                    <ExpansionPanel>
                                        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />} aria-controls="panel1a-content" id="panel1a-header">
                                            <Typography><b>Product Specifications</b> (SKU: {this.state.selectedProductSKU})</Typography>
                                        </ExpansionPanelSummary>
                                        <ExpansionPanelDetails>
                                            <TableContainer component={Paper}>
                                                <Table aria-label="simple table" size="small">
                                                    <TableHead>
                                                    </TableHead>
                                                    <TableBody>
                                                        {this.loadProductFeatures()}
                                                    </TableBody>
                                                </Table>
                                            </TableContainer>
                                        </ExpansionPanelDetails>
                                    </ExpansionPanel>
                                    : null}
                            </Grid>

                        </Grid>

                        <Grid item xs={12} md={12} lg={12} xl={12} className="ProductRating">
                            <ExpansionPanel>
                                <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />} aria-controls="panel1a-content" id="panel1a-header">
                                    <Typography><b>Product Reviews ({this.state.productGroup.CumulativeReviewCount} Reviews)</b></Typography>
                                </ExpansionPanelSummary>
                                <ExpansionPanelDetails>
                                </ExpansionPanelDetails>
                            </ExpansionPanel>
                        </Grid>

                    </Grid>
                }
            </React.Fragment>
        )
    };

}

export default Product;
