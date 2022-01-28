import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import {AppointmentInterface} from "../models/IAppointment";
import { format } from 'date-fns'



const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Appointment() {
  const classes = useStyles();
  const [appointments, setAppointments] = useState<AppointmentInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getAppointments = async () => {
    fetch(`${apiUrl}/appointments`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setAppointments(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getAppointments();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              Appointment List
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/appointment/create"
              variant="contained"
              color="primary"
            >
              Create Data
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="5%">
                  No.
                </TableCell>
                <TableCell align="center" width="25%">
                  Officer
                </TableCell>
                <TableCell align="center" width="25%">
                  Patient
                </TableCell>
                <TableCell align="center" width="25%">
                  Specialist
                </TableCell>
                <TableCell align="center" width="5%">
                  Day(amount)
                </TableCell>
                <TableCell align="center" width="50%">
                  Note
                </TableCell>
                <TableCell align="center" width="50%">
                  AppointDate
                </TableCell>
                <TableCell align="center" width="50%">
                  Date/Time
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {appointments.map((item: AppointmentInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Officer.Name}</TableCell>
                  <TableCell align="center">{item.Patient.Name}</TableCell>
                  <TableCell align="center">{item.Specialist.Name}</TableCell>
                  <TableCell align="center">{item.Number}</TableCell>
                  <TableCell align="center">{item.Note}</TableCell>
                  <TableCell align="center">{format((new Date(item.AppointDate)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                  <TableCell align="center">{format((new Date(item.IssueDate)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Appointment;