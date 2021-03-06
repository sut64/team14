import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import { Select } from "@material-ui/core";
import TextField from '@material-ui/core/TextField';
import DateFnsUtils from '@date-io/date-fns';
import { KeyboardDateTimePicker, MuiPickersUtilsProvider } from "@material-ui/pickers";

import { MedicineandVaccineInterface } from "../models/IMedicineandVaccine";
import { CategoryInterface } from "../models/ICategory";
import { DosageFormInterface } from "../models/IDosageForm";
import { ContagiousInterface } from "../models/IContagious";
import { AgeInterface } from "../models/IAge";

function Alert(props: AlertProps) {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
}
const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: { flexGrow: 1 },
        container: { marginTop: theme.spacing(2) },
        paper: { padding: theme.spacing(2), color: theme.palette.text.secondary },
    })
);
function MedicineandVaccineCreate() {
    const classes = useStyles();
    const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
    const [categories, setCategorys] = useState<CategoryInterface[]>([]);
    const [dosageForms, setDosageForms] = useState<DosageFormInterface[]>([]);
    const [contagious, setContagious] = useState<ContagiousInterface[]>([]);
    const [ages, setAges] = useState<AgeInterface[]>([]);
    const [medicineandvaccines, setMedicineandVaccines] = useState<Partial<MedicineandVaccineInterface>>(
        {}
      );
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const apiUrl = "http://localhost:8080";
    const requestOptions = {
      method: "GET",
      headers: { 
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json" },
    };

    const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };
    
    const handleChange = (
      event: React.ChangeEvent<{ name?: string; value: unknown }>
    ) => {
      const name = event.target.name as keyof typeof medicineandvaccines;
      setMedicineandVaccines({
        ...medicineandvaccines,
        [name]: event.target.value,
      });
    };

      const handleDateChange = (date: Date | null) => {
        console.log(date);
        setSelectedDate(date);
      };

      const getCategory = async () => {
        fetch(`${apiUrl}/categories`, requestOptions)
          .then((response) => response.json())
          .then((res) => {
            if (res.data) {
              setCategorys(res.data);
            } else {
              console.log("else");
            }
          });
      };

      const getDosageForm = async () => {
        fetch(`${apiUrl}/dosage_forms`, requestOptions)
          .then((response) => response.json())
          .then((res) => {
            if (res.data) {
              setDosageForms(res.data);
            } else {
              console.log("else");
            }
          });
      };

      const getContagious = async () => {
        fetch(`${apiUrl}/contagious`, requestOptions)
          .then((response) => response.json())
          .then((res) => {
            if (res.data) {
              setContagious(res.data);
            } else {
              console.log("else");
            }
          });
      };

      const getAge = async () => {
        fetch(`${apiUrl}/ages`, requestOptions)
          .then((response) => response.json())
          .then((res) => {
            if (res.data) {
              setAges(res.data);
            } else {
              console.log("else");
            }
          });
      };

      useEffect(() => {
        getCategory();
        getDosageForm();
        getContagious();
        getAge();
      }, []);

      const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
      };

    function submit() {
        let data = {
                CategoryID: convertType(medicineandvaccines.CategoryID),
                DosageFormID: convertType(medicineandvaccines.DosageFormID),
                AgeID: convertType(medicineandvaccines.AgeID),
                ContagiousID: convertType(medicineandvaccines.ContagiousID),
                RegNo: medicineandvaccines.RegNo,
                Name: medicineandvaccines.Name,
                Date: selectedDate,
                MinAge: Number(medicineandvaccines.MinAge),
                MaxAge: Number(medicineandvaccines.MaxAge),
        };
        console.log(data)

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/medicineand_vaccines`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("???????????????????????????")
          setSuccess(true);
          setErrorMessage("");
        } else {
          console.log("????????????????????????????????????")
          setError(true);
          setErrorMessage(res.error);
        }
      });
  }
    return (
        <div>
        <Container className={classes.container} maxWidth="md">
            <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="success">
                    ??????????????????????????????????????????????????????
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="error">
                    ???????????????????????????????????????????????????????????????: {errorMessage}
                </Alert>
            </Snackbar>
            <Paper className={classes.paper}>
                <Box display="flex">
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            ???????????????????????????????????????????????????
                        </Typography>
                    </Box>
                </Box>
                <Divider />
            
                <Grid container spacing={3} className={classes.root}>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>??????????????????????????????</p>
                            <TextField
                                name="RegNo"
                                type="string"
                                multiline
                                rows={1}
                                value={medicineandvaccines.RegNo}
                                onChange={handleChange}
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>????????????</p>
                            <TextField
                                name="Name"
                                type="string"
                                multiline
                                rows={1}
                                value={medicineandvaccines.Name}
                                onChange={handleChange}
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>??????????????????</p>
                            <Select
                                native
                                value={medicineandvaccines.CategoryID}
                                onChange={handleChange}
                                inputProps={{
                                name: "CategoryID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    ????????????????????????????????????????????????
                                </option>
                                   {categories.map((item: CategoryInterface) => (
                                <option value={item.ID} key={item.ID}>
                                    {item.Category}
                                </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>????????????????????????</p>
                            <Select
                                native
                                value={medicineandvaccines.DosageFormID}
                                onChange={handleChange}
                                inputProps={{
                                name: "DosageFormID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    ??????????????????????????????????????????????????????
                                </option>
                                   {dosageForms.map((item: DosageFormInterface) => (
                                <option value={item.ID} key={item.ID}>
                                    {item.DosageForm}
                                </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                      <FormControl fullWidth variant="outlined">
                      <p>MinAge</p>
                        <TextField
                          name="MinAge"
                          variant="outlined"
                          type="number"
                          size="medium"
                          placeholder="???????????????????????????????????????"
                          value={medicineandvaccines.MinAge}
                          onChange={handleChange}
                        />    
                      </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                      <FormControl fullWidth variant="outlined">
                      <p>MaxAge</p>
                        <TextField
                          name="MaxAge"
                          variant="outlined"
                          type="number"
                          size="medium"
                          placeholder="???????????????????????????????????????"
                          value={medicineandvaccines.MaxAge}
                          onChange={handleChange}
                        />  
                      </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>????????????????????????</p>
                            <Select
                                native
                                value={medicineandvaccines.AgeID}
                                onChange={handleChange}
                                inputProps={{
                                name: "AgeID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    ??????????????????????????????????????????????????????
                                </option>
                                   {ages.map((item: AgeInterface) => (
                                <option value={item.ID} key={item.ID}>
                                    {item.Age}
                                </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>????????????????????????</p>
                            <Select
                                native
                                value={medicineandvaccines.ContagiousID}
                                onChange={handleChange}
                                inputProps={{
                                name: "ContagiousID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    ?????????????????????????????????????????????????????????
                                </option>
                                   {contagious.map((item: ContagiousInterface) => (
                                <option value={item.ID} key={item.ID}>
                                    {item.Name}
                                </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>????????????????????????????????????</p>
                            <MuiPickersUtilsProvider utils={DateFnsUtils}>
                                <KeyboardDateTimePicker
                                    name="Date"
                                    value={selectedDate}
                                    onChange={handleDateChange}
                                    label="?????????????????????????????????????????????????????????????????????"
                                    minDate={new Date("2020-01-01T00:00")}
                                    format="yyyy/MM/dd hh:mm a"
                                />
                            </MuiPickersUtilsProvider>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <Button
                            component={RouterLink}
                            to="/medicineand_vaccines"
                            variant="contained"
                        >
                            ????????????
                        </Button>
                        <Button
                            style={{ float: "right" }}
                            variant="contained"
                            onClick={submit}
                            color="primary"
                        >
                            ??????????????????
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
        </div>
    )
};
export default MedicineandVaccineCreate;

