
        // onload为浏览器对象中的事件,页面载入时触发 
        window.onload = function() {
            var grid;
            var chessArr = [];
            var timer = 0;
            var lineNum = 15;
            // 获取元素 
            var box = document.getElementById('chessboard');
            var chessBox = box.getElementsByTagName('div');
            var submitBtn = document.getElementById('submitBtn');
            submitBtn.onclick = function() {
                    var chessMaxNum = lineNum * lineNum;
                    var chessWH = 90 / lineNum;
                    for (var i = 0; i < chessMaxNum; i++) {
                        // 设置棋盘里小格子div元素 
                        grid = document.createElement('div');
                        grid.style.width = 'calc(' + chessWH + 'vmin - 2px)';
                        grid.style.height = 'calc(' + chessWH + 'vmin - 2px)';
                        grid.id = i;
                        box.appendChild(grid);
                        chessArr[i] = 0;
                        grid.onclick = function(x) {
                            // target 事件属性可返回事件的目标节点（触发该事件的节点），如生成事件的元素、文档或窗口。 
                            var index = x.target.id || x.target.parentNode.id;
                            return playChess(index);
                        };
                    };
                    mask.style.display = 'none';
                }
                //棋子对象 
            function Chess() {
                this.color = 'white';
                this.site = 0;
                // 创建一个class 
                this.chessDom = function() {
                    // 创造新节点 
                    var dom = document.createElement('p');
                    // 将这个名字给class 
                    dom.setAttribute('class', this.color);
                    return dom;
                }
                this.ligature = function(arr) {
                    // map() 方法返回一个新数组，数组中的元素为原始数组元素调用函数处理后的值。 
                    //是就返回这个site 
                    // 给白棋一个标识号,方便在下列judge中判断 
                    var whiteChess = arr.map(function(s) {
                        // parseInt() 函数可解析一个字符串，并返回一个整数。 
                        return (s.color == 'white') ? parseInt(s.site) : 0;
                    });
                    var blackChess = arr.map(function(s) {
                        return (s.color == 'black') ? parseInt(s.site) : 0;
                    });
                }
            }

            function playChess(i) {
                if (chessArr[i] == 0) {
                    chessArr[i] = new Chess();
                    chessArr[i].color = 'white';
                    chessBox[i].appendChild(chessArr[i].chessDom());


                    // ai
                    console.log("ai's turn")
                    chessArr[i].ligature(chessArr);
                    chessArr[parseInt(i)+1] = new Chess();
                    chessArr[parseInt(i)+1].color = 'black';
                    chessBox[parseInt(i)+1].appendChild(chessArr[parseInt(i)+1].chessDom());
                    chessArr[parseInt(i)+1].ligature(chessArr);
                } else {
                    alert('此处有子!');
                }
            }
        };