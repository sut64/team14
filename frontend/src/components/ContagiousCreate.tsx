import React from "react";
import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import DateFnsUtils from '@date-io/date-fns';
import {
  MuiPickersUtilsProvider,
  KeyboardDatePicker,
} from '@material-ui/pickers';

import { GermInterface } from "../models/IGerm";
import { CatchingTypeInterface } from "../models/ICatchingType";
import { RiskGroupTypeInterface } from "../models/IRiskGroupType";
import { ContagiousInterface } from "../models/IContagious";


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

function ContagiousCreate() {
  const classes = useStyles();
  const [germ, setGerm] = useState<GermInterface[]>([]);
  const [catching_type, setCatchingType] = useState<CatchingTypeInterface[]>([]);
  const [risk_group_type, setRiskGroupType] = useState<RiskGroupTypeInterface[]>([]);
  const [contagious, setContagious] = useState<Partial<ContagiousInterface>>({});
  const [selectedDate, setSelectedDate] = React.useState<Date | null>(
    new Date()
  );

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
    const name = event.target.name as keyof typeof contagious;
    setContagious({
      ...contagious,
      [name]: event.target.value,
    });
  };


  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof contagious;
    const { value } = event.target;
    setContagious({ ...contagious, [id]: value });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };


  const getGerm = async () => {
    fetch(`${apiUrl}/germ`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setGerm(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getCatchingType = async () => {
    fetch(`${apiUrl}/catching_type`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setCatchingType(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getRiskGroupType = async () => {
    fetch(`${apiUrl}/risk_group_type`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setRiskGroupType(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getGerm();
    getCatchingType();
    getRiskGroupType();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      GermID: convertType(contagious.GermID),
      CatchingTypeID: convertType(contagious.CatchingTypeID),
      RiskGroupTypeID: convertType(contagious.RiskGroupTypeID),
      Name: contagious.Name ?? "",
      Symptom: contagious.Symptom ?? "",
      Incubation: typeof contagious.Incubation === "string" ? parseInt(contagious.Incubation) : 0,
      Date: selectedDate,
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

    fetch(`${apiUrl}/contagious`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true)
          setErrorMessage("")
        } else {
          console.log("บันทึกไม่ได้")
          setError(true)
          setErrorMessage(res.error)
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
              บันทึกข้อมูลโรคติดต่อ
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={2} className={classes.root}>
          <Grid item xs={12}>
            <p>ชื่อโรคติดต่อ</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Name"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกชื่อโรคติดต่อ"
                value={contagious.Name || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>  
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เชื้อโรคที่เป็นสาเหตุ</p>
              <Select
                native
                value={contagious.GermID}
                onChange={handleChange}
                inputProps={{
                  name: "GermID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกเชื้อโรค
                </option>
                {germ.map((item: GermInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ประเภทการติดต่อ</p>
              <Select
                native
                value={contagious.CatchingTypeID}
                onChange={handleChange}
                inputProps={{
                  name: "CatchingTypeID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกประเภทการติดต่อ
                </option>
                {catching_type.map((item: CatchingTypeInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Title}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
              <p>อาการ</p>
              <TextField
                id="Symptom"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกข้อมูลอาการ"
                value={contagious.Symptom || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={3}>
           <FormControl fullWidth variant="outlined">
             <p>ระยะฟักตัวโดยเฉลี่ย (วัน)</p>
             <TextField
               id="Incubation"
               variant="outlined"
               type="number"
               size="medium"
               InputProps={{ inputProps: { min: 1 } }}
               InputLabelProps={{
                 shrink: true,
               }}
               value={contagious.Incubation || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
         <Grid item xs={5}>
            <FormControl fullWidth variant="outlined">
              <p>ประเภทกลุ่มเสี่ยง</p>
              <Select
                native
                value={contagious.RiskGroupTypeID}
                onChange={handleChange}
                inputProps={{
                  name: "RiskGroupTypeID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกประเภทกลุ่มเสี่ยง
                </option>
                {risk_group_type.map((item: RiskGroupTypeInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Title}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={4}>
           <FormControl fullWidth variant="outlined">
             <p>วันที่บันทึก</p>
             <MuiPickersUtilsProvider utils={DateFnsUtils}>
             <KeyboardDatePicker
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
          
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/contagious"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default ContagiousCreate;
