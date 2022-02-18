import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

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

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบจัดการโรคติดต่อ</h1>
        <h4>about</h4>
        <p>
        ระบบจัดการโรคติดต่อของโรงพยาบาล เป็นระบบที่มีไว้สำหรับเจ้าหน้าที่ในแผนกโรคติดต่อ เป็นระบบที่สามารถบันทึกข้อมูลเกี่ยวกับข้อมูลต่างๆเกี่ยวกับการจัดการโรคติดต่อภายในโรงพยาบาล 
        ประกอบไปด้วย 
        </p>
        
        <h4>ระบบบันทึกการเข้าใช้ห้อง </h4> 
        <h4>ระบบบันทึกข้อมูลยาเเละวัคซีน</h4>
        <h4>ระบบบันทึกการติดตามผลการรักษา</h4>
        <h4>ระบบบันทึกข้อมูลโรคติดต่อ</h4>
        <h4>ระบบคัดกรองผู้ป่วย</h4>
        <h4>ระบบบันทึกการป้องกันโรคติดต่อ</h4>


      </Container>
    </div>
  );
}
export default Home;