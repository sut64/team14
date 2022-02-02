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

import { AppointmentInterface } from "../models/IAppointment";
import { OfficersInterface } from "../models/IOfficer";
import { SpecialistsInterface } from "../models/ISpecialist";
import { PatientsInterface } from "../models/IPatient";

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
function AppointmentCreate() {
    const classes = useStyles();
    const [appointDate, setAppointDate] = useState<Date | null>(new Date());
    const [issueDate, setIssueDate] = useState<Date | null>(new Date());
    const [officers, setOfficers] = useState<OfficersInterface>();
    const [specialists, setSpecialists] = useState<SpecialistsInterface[]>([]);
    const [patients, setPatients] = useState<PatientsInterface[]>([]);
    const [appointment, setAppointment] = useState<Partial<AppointmentInterface>>(
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
    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: unknown }>
      ) => {
        const name = event.target.name as keyof typeof appointment;
        setAppointment({
          ...appointment,
          [name]: event.target.value,
        });
      };

    const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };
     
    const handleInputChange = (
        event: React.ChangeEvent<{ name?: string; value: any }>
      ) => {
        const name = event.target.name as keyof typeof appointment;
        const { value } = event.target;
        setAppointment({ 
            ...appointment, 
            [name]: value });
      };
    
      const handleIssueDateChange = (date: Date | null) => {
        console.log(date);
        setIssueDate(date);
      };
      const handleAppointDateChange = (date: Date | null) => {
        console.log(date);
        setAppointDate(date);
      };

      const getOfficer = async () => {
        let uid = localStorage.getItem("uid");
        fetch(`${apiUrl}/officer/ ${uid}`, requestOptions)
          .then((response) => response.json())
          .then((res) => {
            appointment.OfficerID = res.data.ID
            if (res.data) {
              setOfficers(res.data);
            } else {
              console.log("else");
            }
          });
      };

      const getSpecialists = async () => {
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

      const getPatients = async () => {
        fetch(`${apiUrl}/patients`, requestOptions)
          .then((response) => response.json())
          .then((res) => {
            if (res.data) {
              setPatients(res.data);
            } else {
              console.log("else");
            }
          });
      };

      useEffect(() => {
        getOfficer();
        getSpecialists();
        getPatients();
      }, []);

      const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
      };

    function submit() {
        let data = {
                OfficerID: convertType(appointment.OfficerID),
                SpecialistID: convertType(appointment.SpecialistID),
                PatientID: convertType(appointment.PatientID),
                Note: appointment.Note,
                IssueDate: issueDate,
                AppointDate: appointDate,
                Number: convertType(appointment.Number) ,
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

    fetch(`${apiUrl}/appointments`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
          setErrorMessage("");
        } else {
          console.log("บันทึกไม่ได้",res)
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
                    บันทึกข้อมูลสำเร็จ 
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="error">
                    บันทึกข้อมูลไม่สำเร็จ {errorMessage}
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
                            บันทึกการนัดตรวจโรค
                        </Typography>
                    </Box>
                </Box>
                <Divider />
            
                <Grid container spacing={3} className={classes.root}>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>เจ้าหน้าที่</p>
                            <Select
                                  native
                                  value={appointment.OfficerID}
                                  onChange={handleChange}
                                  disabled
                                  inputProps={{
                                    name: "OfficerID",
                                  }}
                                >
                                  <option aria-label="None" value="">
                                    กรุณาเลือกรายการเจ้าหน้าที่
                                  </option>
                                  <option value={officers?.ID} key={officers?.ID}>
                                    {officers?.Name}
                                  </option>
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>รายการผู้ป่วย</p>
                            <Select
                                native
                                value={appointment.PatientID}
                                onChange={handleChange}
                                inputProps={{
                                name: "PatientID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกรายการผู้ป่วย
                                </option>
                                   {patients.map((item: PatientsInterface) => (
                                <option value={item.ID} key={item.ID}>
                                    {item.Name}
                                </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>รายการแพทย์</p>
                            <Select
                                native
                                value={appointment.SpecialistID}
                                onChange={handleChange}
                                inputProps={{
                                name: "SpecialistID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกรายการแพทย์
                                </option>
                                   {specialists.map((item: SpecialistsInterface) => (
                                <option value={item.ID} key={item.ID}>
                                    {item.Name}
                                </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>หมายเหตุ</p>
                            <TextField
                                name="Number"
                                type="string"
                                multiline
                                rows={1}
                                value={appointment.Number}
                                onChange={handleChange}
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>หมายเหตุ</p>
                            <TextField
                                name="Note"
                                type="string"
                                multiline
                                rows={1}
                                value={appointment.Note}
                                onChange={handleChange}
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>วันที่นัด</p>
                            <MuiPickersUtilsProvider utils={DateFnsUtils}>
                                <KeyboardDateTimePicker
                                    name="Date"
                                    value={appointDate}
                                    onChange={handleAppointDateChange}
                                    label="กรุณาเลือกวันที่และเวลา"
                                    minDate={new Date("2020-01-01T00:00")}
                                    format="yyyy/MM/dd hh:mm a"
                                />
                            </MuiPickersUtilsProvider>
                        </FormControl>
                    </Grid>

                  
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>วันที่บันทึก</p>
                            <MuiPickersUtilsProvider utils={DateFnsUtils}>
                                <KeyboardDateTimePicker
                                    name="Date"
                                    value={issueDate}
                                    onChange={handleIssueDateChange}
                                    label="กรุณาเลือกวันที่และเวลา"
                                    minDate={new Date("2020-01-01T00:00")}
                                    format="yyyy/MM/dd hh:mm a"
                                />
                            </MuiPickersUtilsProvider>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <Button
                            component={RouterLink}
                            to="/appointment"
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
        </div>
    )
};
export default AppointmentCreate;