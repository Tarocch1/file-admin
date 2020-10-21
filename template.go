package main

var listTemplate = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>{{ .Title }}</title>
	<style>
	table, th, td {
		padding: 5px;
		border: 1px solid black;
	}
	table {
		border-collapse: collapse;
		border-spacing: 0px;
	}
	svg {
		display: block;
		margin: auto;
	}
	</style>
</head>
<body>
	<table>
		<thead>
			<tr>
				<th colspan="5">{{ .Title }}</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td style="width: 16px;"><svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" width="16" height="16"><path d="M853.333333 245.333333H245.333333l93.866667-93.866666c12.8-12.8 12.8-34.133333 0-46.933334-12.8-12.8-34.133333-12.8-46.933333 0l-145.066667 145.066667c-12.8 12.8-12.8 34.133333 0 46.933333l145.066667 145.066667c6.4 6.4 14.933333 10.666667 23.466666 10.666667s17.066667-4.266667 23.466667-10.666667c12.8-12.8 12.8-34.133333 0-46.933333L256 311.466667h597.333333c6.4 0 10.666667 4.266667 10.666667 10.666666v426.666667c0 6.4-4.266667 10.666667-10.666667 10.666667H170.666667c-17.066667 0-32 14.933333-32 32s14.933333 32 32 32h682.666666c40.533333 0 74.666667-34.133333 74.666667-74.666667V320c0-40.533333-34.133333-74.666667-74.666667-74.666667z"></path></svg></td>
				<td colspan="4"><a href="../">..</a></td>
			</tr>
			{{ range .Files }}
			<tr>
				{{ if .IsDir }}
				<td style="width: 16px;"><svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" width="16" height="16"><path d="M853.333333 266.666667H514.133333c-4.266667 0-6.4-2.133333-8.533333-4.266667l-38.4-66.133333c-12.8-21.333333-38.4-36.266667-64-36.266667H170.666667c-40.533333 0-74.666667 34.133333-74.666667 74.666667v554.666666c0 40.533333 34.133333 74.666667 74.666667 74.666667h682.666666c40.533333 0 74.666667-34.133333 74.666667-74.666667V341.333333c0-40.533333-34.133333-74.666667-74.666667-74.666666z m-682.666666-42.666667h232.533333c4.266667 0 6.4 2.133333 8.533333 4.266667l38.4 66.133333c12.8 21.333333 38.4 36.266667 64 36.266667H853.333333c6.4 0 10.666667 4.266667 10.666667 10.666666v74.666667h-704V234.666667c0-6.4 4.266667-10.666667 10.666667-10.666667z m682.666666 576H170.666667c-6.4 0-10.666667-4.266667-10.666667-10.666667V480h704V789.333333c0 6.4-4.266667 10.666667-10.666667 10.666667z"></path></svg></td>
				{{ else }}
				<td style="width: 16px;"><svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" width="16" height="16"><path d="M842.666667 285.866667l-187.733334-187.733334c-14.933333-14.933333-32-21.333333-53.333333-21.333333H234.666667C194.133333 74.666667 160 108.8 160 149.333333v725.333334c0 40.533333 34.133333 74.666667 74.666667 74.666666h554.666666c40.533333 0 74.666667-34.133333 74.666667-74.666666V337.066667c0-19.2-8.533333-38.4-21.333333-51.2z m-44.8 44.8c-2.133333 2.133333-4.266667 0-8.533334 0h-170.666666c-6.4 0-10.666667-4.266667-10.666667-10.666667V149.333333c0-2.133333 0-6.4-2.133333-8.533333 0 0 2.133333 0 2.133333 2.133333l189.866667 187.733334z m-8.533334 554.666666H234.666667c-6.4 0-10.666667-4.266667-10.666667-10.666666V149.333333c0-6.4 4.266667-10.666667 10.666667-10.666666h311.466666c-2.133333 4.266667-2.133333 6.4-2.133333 10.666666v170.666667c0 40.533333 34.133333 74.666667 74.666667 74.666667h170.666666c4.266667 0 6.4 0 10.666667-2.133334V874.666667c0 6.4-4.266667 10.666667-10.666667 10.666666z"></path><path d="M640 693.333333H341.333333c-17.066667 0-32 14.933333-32 32s14.933333 32 32 32h298.666667c17.066667 0 32-14.933333 32-32s-14.933333-32-32-32zM640 522.666667H341.333333c-17.066667 0-32 14.933333-32 32s14.933333 32 32 32h298.666667c17.066667 0 32-14.933333 32-32s-14.933333-32-32-32zM341.333333 416h85.333334c17.066667 0 32-14.933333 32-32s-14.933333-32-32-32h-85.333334c-17.066667 0-32 14.933333-32 32s14.933333 32 32 32z"></path></svg></td>
				{{ end }}
				<td><a href="{{ .Name }}{{ if .IsDir }}/{{ end }}">{{ .Name }}</a></td>
				<td data-mod-time="{{ .Time }}" />
				<td style="text-align: end;">{{ formatSize .Size }}</td>
				<td><button onclick="deleteFile({{ .Name }});">删除</button></td>
			</tr>
			{{ end }}
		</tbody>
		<tfoot>
			<tr>
				<td colspan="5">
					<button onclick="openFileSelector()">上传</button>
					<button onclick="createFolder()">新建文件夹</button>
				</th>
			</tr>
		</tfoot>
	</table>
	<input type="file" id="input_file" style="display:none" onchange="uploadFiles(this.files)" multiple>
	<script>
	const timeFormatter = Intl.DateTimeFormat(undefined, {
		year: 'numeric',
		month: '2-digit',
		day: '2-digit',
		hour: '2-digit',
		minute: '2-digit',
		second: '2-digit',
		hourCycle: 'h23',
	});
	window.onload = function () {
		document
			.querySelectorAll('td[data-mod-time]')
			.forEach(
				td =>
					(td.innerText = timeFormatter.format(
						new Date(Number(td.dataset['modTime']) * 1000),
					)),
			);
	};
	window.ondragenter = function (e) {
		e.stopPropagation();
		e.preventDefault();
	};
	window.ondragover = function (e) {
		e.stopPropagation();
		e.preventDefault();
	};
	window.ondragleave = function (e) {
		e.stopPropagation();
		e.preventDefault();
	};
	window.ondrop = function (e) {
		e.stopPropagation();
		e.preventDefault();
		uploadFiles(e.dataTransfer.files);
	};
	function openFileSelector() {
		document.querySelector('#input_file').click();
	}
	function uploadFiles(files) {
		Promise.allSettled(
			[...files].map(file =>
				fetch(file.name, {
					method: 'PUT',
					body: file,
				}),
			),
		).then(results => {
			const all = results.length;
			const succeed = results.filter(
				result =>
					result.status === 'fulfilled' &&
					(result.value.status === 200 || result.value.status === 201),
			).length;
			alert(succeed + ' / ' + all + ' 个文件上传完成。');
			window.location.reload();
		});
	}
	function createFolder() {
		const folderName = prompt('请输入文件夹名：');
		if (folderName) {
			fetch(folderName, {
				method: 'POST',
				headers: {
					'X-Is-Dir': 'true',
				},
			}).then(res => {
				if (res.status === 201) {
					window.location.reload();
				} else {
					alert('新建文件夹失败！');
				}
			});
		}
	}
	function deleteFile(path) {
		const r = confirm('是否要删除 ' + path + ' ？');
		if (r) {
			fetch(path, {
				method: 'DELETE',
			}).then(res => {
				if (res.status === 200) {
					window.location.reload();
				} else {
					alert('删除失败！');
				}
			});
		}
	}
	</script>
</body>
</html>
`
