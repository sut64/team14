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

import { OfficersInterface } from "../models/IOfficer";
import { ContagiousInterface } from "../models/IContagious";
import { SpecialistsInterface } from "../models/ISpecialist";
import { PreventionsInterface } from "../models/IPrevention";



const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function PreventionCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  
  const [Officer, setOfficers] = useState<OfficersInterface[]>([]);
  const [Contagious, setContagious] = useState<ContagiousInterface[]>([]);
  const [Specialist, setSpecialists] = useState<SpecialistsInterface[]>([]);
  const [Prevention, setPreventions] = useState<Partial<PreventionsInterface>>( {});
 

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
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
    const name = event.target.name as keyof typeof Prevention;
    setPreventions({
      ...Prevention,
      [name]: event.target.value,
    });
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof Prevention;
    const { value } = event.target;
    setPreventions({ ...Prevention, [id]: value });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getOfficer = async () => {
    fetch(`${apiUrl}/officers`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setOfficers(res.data);
        } else {
          console.log("else");
        }
      });
  };


  const getSpecialist = async () => {
    fetch(`${apiUrl}/specialists`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSpecialists(res.data);
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

  

  useEffect(() => {
    getOfficer();
    getContagious();
    getSpecialist();
   
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      OfficerID: convertType(Prevention.OfficerID),
      ContagiousID: convertType(Prevention.ContagiousID),
      Disease: Prevention.Disease,
      SpecialistID: convertType(Prevention.SpecialistID),
      Protection: Prevention.Protection,
      Date: selectedDate,
      Age: typeof Prevention.Age === "string" ? parseInt(Prevention.Age) : 0 ,
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

    fetch(`${apiUrl}/preventions`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
          setErrorMessage("");
        } else {
          console.log("บันทึกไม่ได้")
          setError(true);
          setErrorMessage(res.error);
        }
      });
  }
  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ: {errorMessage}
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
              บันทึกข้อมูลการป้องกันโรค
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เลือกเจ้าหน้าที่แผนกโรคติดต่อ</p>
              <Select
                native
                value={Prevention.OfficerID}
                onChange={handleChange}
                inputProps={{
                  name: "OfficerID",
                }}
              >
                <option aria-label="None" value="">
                  โปรดเลือกเจ้าหน้าที่แผนกโรคติดต่อ
                </option>
                {Officer.map((item: OfficersInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เลือกข้อมูลโรคติดต่อ</p>
              <Select
                native
                value={Prevention.ContagiousID}
                onChange={handleChange}
                inputProps={{
                  name: "ContagiousID",
                }}
              >
                <option aria-label="None" value="">
                  โปรดเลือกข้อมูลโรคติดต่อ
                </option>
                {Contagious.map((item: ContagiousInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>กรอกข้อมูลโรค</p>
              <TextField
                name="Disease"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกข้อมูลโรค"
                value={Prevention.Disease}
                onChange={handleChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เลือกแพทย์ผู้ให้ข้อมูลเฉพาะทาง</p>
              <Select
                native
                value={Prevention.SpecialistID}
                onChange={handleChange}
                inputProps={{
                  name: "SpecialistID",
                }}
              >
                <option aria-label="None" value="">
                  โปรดเลือกแพทย์ผู้ให้ข้อมูลเฉพาะทาง
                </option>
                {Specialist.map((item: SpecialistsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>กรอกวิธีการป้องกันโรคติดต่อ</p>
              <TextField
                name="Protection"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกวิธีการป้องกันโรคติดต่อ"
                value={Prevention.Protection}
                onChange={handleChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                    <p>วันที่บันทึก</p>
                <MuiPickersUtilsProvider utils={DateFnsUtils}>
                    <KeyboardDateTimePicker
                        margin="normal"
                        id="date-picker-dialog"
                        label=""
                        format="dd/MM/yyyy"
                        value={selectedDate}
                        onChange={handleDateChange}
                        KeyboardButtonProps={{
                          'aria-label': 'change date',
                        }}
                    />
                </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>กรอกอายุแพทย์</p>
              <TextField
                name="Age"
                variant="outlined"
                type="int"
                size="medium"
                placeholder="โปรดกรอกอายุแพทย์"
                value={Prevention.Age}
                onChange={handleChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/preventions"
              variant="contained"
            >
              BACK
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              SUBMIT
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}
export default PreventionCreate;