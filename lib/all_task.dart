
import 'dart:convert';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:task_manager2/add_task_page.dart';

class AllTask extends StatefulWidget {


  const AllTask({Key? key}) : super(key: key);



  @override
  State<AllTask> createState() => _AllTaskState();
}

class _AllTaskState extends State<AllTask> {

  /////////////////////////////////////

  Map <dynamic, dynamic>? apiMap;
  List<dynamic>? list;


  Future getUserData() async
  {
    var response=await http.get(
        Uri.parse('http://192.168.0.204:8080/gettasks')
    );

    setState((){
      list = jsonDecode(response.body);
      // list=list?..sort((a,b)=>a['id'].compareTo(b['id']));    [index]['created_at'].toString()
      list=list?..sort((a,b)=>a['created_at'].compareTo(b['created_at']));

    });

    print(list?.length);



  }

  // Future getUserData() async
  // {
  //   var response=await http.get(
  //       Uri.parse('https://herosapp.nyc3.digitaloceanspaces.com/quiz.json')
  //   );
  //
  //
  //   setState((){
  //     apiMap = jsonDecode(response.body);
  //   });
  //  // print(list?.length);
  // }







  ///////////////////////////////////




  @override
  Widget build(BuildContext context) {
    return list!=null?MaterialApp(
      home: Scaffold(
        body: ListView.builder(
            itemCount: list!.length,

            itemBuilder:(BuildContext contex,int index){

          return ListTile(
            title: Text(list![index]["task_name"]),
            subtitle: Text(list![index]["task_detail"]),

          );
        }),

        floatingActionButtonLocation: FloatingActionButtonLocation.endDocked,
        floatingActionButton: Container(
          margin: EdgeInsets.fromLTRB(0, 0, 0, 21),
          child: FloatingActionButton.extended(

            onPressed: (){

              Navigator.push(
                  context,
                  MaterialPageRoute(
                      builder: (context) =>
                          AddTaskPage()));

            },
            label: Text("Add Task"),
            hoverElevation: 100,
            icon: Icon(Icons.edit),
            splashColor: Colors.purple,

            backgroundColor: Colors.pink.shade500,
          ),
        ),





      ),
    ):Scaffold(
        body: Center(
          child: Text("Loading-----------",
            style: TextStyle(
                fontSize: 30,
                fontWeight: FontWeight.bold
            ),


          ),
        )
    );}

  @override
  void initState() {

    getUserData();
    super.initState();

  }
}
