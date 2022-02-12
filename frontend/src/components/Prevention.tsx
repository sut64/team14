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
import { PreventionsInterface } from "../models/IPrevention";
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

function Preventions() {
  const classes = useStyles();
  const [preventions, setPreventions] = useState<PreventionsInterface[]>([]);

  const apiUrl = "http://localhost:8080";
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };
  
    const getPreventions = async () => {
    

    fetch(`${apiUrl}/preventions` , requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setPreventions(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getPreventions();
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
              Prevention Data
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/prevention/create"
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
                <TableCell align="center" width="10%">
                  No.
                </TableCell>
                <TableCell align="center" width="25%">
                  Officer
                </TableCell>
                <TableCell align="center" width="25%">
                  Contagious
                </TableCell>
                <TableCell align="center" width="25%">
                  Disease
                </TableCell>
                <TableCell align="center" width="25%">
                  Specialist
                </TableCell>
                <TableCell align="center" width="25%">
                  Protection
                </TableCell>
                <TableCell align="center" width="25%">
                  Age
                </TableCell>
                <TableCell align="center" width="50%">
                  Date
                </TableCell>
                
              </TableRow>
            </TableHead>
            <TableBody>
              {preventions.map((item: PreventionsInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>                
                  <TableCell align="center">{item.Officer.Name}</TableCell>
                  <TableCell align="center">{item.Contagious.Name}</TableCell>
                  <TableCell align="center">{item.Disease}</TableCell>
                  <TableCell align="center">{item.Specialist.Name}</TableCell>                  
                  <TableCell align="center">{item.Protection}</TableCell>                
                  <TableCell align="center">{item.Age}</TableCell>
                  <TableCell align="center">{format((new Date(item.Date)), 'dd MM yyyy')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Preventions;